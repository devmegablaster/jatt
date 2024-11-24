package reader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/styles"
)

// TODO: Improve error handling and logging

type Reader struct {
	Cfg config.JattConfig
}

type ListingItem struct {
	Name        string
	Title       string
	Description string
	URL         string
	Date        string
}

type File struct {
	Name        string
	Content     []byte
	FrontMatter map[string]string
	Listing     []ListingItem
}

type FrontMatter struct {
	Layout      string   `yaml:"layout"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Date        string   `yaml:"date"`
	Tags        []string `yaml:"tags"`
	Draft       bool     `yaml:"draft"`
}

func NewReader(cfg config.JattConfig) *Reader {
	return &Reader{
		Cfg: cfg,
	}
}

func (r *Reader) Read() []File {
	files := r.ReadDir(r.Cfg.SiteConfig.ContentDir)
	return files
}

func (r *Reader) ReadDir(dirPath string) []File {
	files := []File{}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			fmt.Println(styles.ErrorStyle.Render("Directory " + dirPath + " does not exist"))
		default:
			fmt.Println(styles.ErrorStyle.Render("Error reading directory " + dirPath))
		}
		os.Exit(1)
	}

	for _, file := range dir {
		if file.IsDir() {
			dirFiles := r.ReadDir(dirPath + "/" + file.Name())
			for _, dirFile := range dirFiles {
				files = append(files, dirFile)
			}
			continue
		}

		f, err := os.Open(dirPath + "/" + file.Name())
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error reading file " + file.Name()))
			os.Exit(1)
		}
		defer f.Close()

		buf := bytes.Buffer{}
		if _, err := io.Copy(&buf, f); err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error reading file " + file.Name()))
			os.Exit(1)
		}

		fm := r.ReadFrontMatter(buf.Bytes())

		fileStruct := File{
			Name:        strings.Replace(dirPath+"/"+strings.Split(file.Name(), ".")[0], "content/", "", 1),
			Content:     []byte(strings.Join(strings.Split(string(buf.Bytes()), "---")[2:], "")),
			FrontMatter: fm,
		}

		if fm["layout"] == "listing" {
			listingItem := r.GenerateListing(fileStruct)
			fileStruct.Listing = listingItem
		}

		files = append(files, fileStruct)

		fmt.Println(styles.DebugStyle.Render("Read " + file.Name()))
	}

	return files
}

func (r *Reader) ReadFrontMatter(buf []byte) map[string]string {
	frontMatter := make(map[string]string)

	if buf[0] != '-' || buf[1] != '-' || buf[2] != '-' {
		return frontMatter
	}

	buf = buf[4:]
	for {
		if buf[0] == '-' && buf[1] == '-' && buf[2] == '-' {
			break
		}

		key := []byte{}
		for {
			if buf[0] == ':' {
				break
			}
			key = append(key, buf[0])
			buf = buf[1:]
		}

		buf = buf[1:]
		buf = buf[1:]
		value := []byte{}
		for {
			if buf[0] == '\n' {
				break
			}
			value = append(value, buf[0])
			buf = buf[1:]
		}

		frontMatter[string(key)] = string(value)

		buf = buf[1:]
	}

	return frontMatter
}

func (r *Reader) GenerateListing(fileStruct File) []ListingItem {
	dirPath := strings.Join(strings.Split(fileStruct.Name, "/")[0:len(strings.Split(fileStruct.Name, "/"))-1], "/")
	dirPath = r.Cfg.SiteConfig.ContentDir + "/" + dirPath

	files := []ListingItem{}

	dir, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error reading directory " + dirPath))
		os.Exit(1)
	}

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if strings.Split(file.Name(), ".")[0] == "index" {
			continue
		}

		f, err := os.Open(dirPath + "/" + file.Name())
		if err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error reading file " + file.Name()))
			os.Exit(1)
		}
		defer f.Close()

		buf := make([]byte, 1024)
		for {
			n, err := f.Read(buf)
			if err != nil && err != io.EOF {
				fmt.Println(styles.ErrorStyle.Render("Error reading file " + file.Name()))
				os.Exit(1)
			}
			if n == 0 {
				break
			}
		}

		fm := r.ReadFrontMatter(buf.Bytes())
		if fm.Draft {
			continue
		}
		files = append(files, ListingItem{
			Name:        strings.Split(file.Name(), ".")[0],
			URL:         strings.Replace(dirPath+"/"+strings.Split(file.Name(), ".")[0], "content", "", 1),
			Date:        fm["date"],
			Title:       fm["title"],
			Description: fm["description"],
		})
	}

	fmt.Println(styles.DebugStyle.Render("Generated listing for " + fileStruct.Name))

	fileStruct.Listing = files
	return files
}
