package reader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/devmegablaster/jatt/internal/config"
	"github.com/devmegablaster/jatt/pkg/styles"
	"gopkg.in/yaml.v3"
)

// TODO: Improve error handling and logging

type Reader struct {
	Cfg config.JattConfig
}

type ListingItem struct {
	ID          string
	Name        string
	Title       string `xml:"title"`
	Description string `xml:"description"`
	URL         string `xml:"link"`
	Date        string `xml:"pubDate"`
	Tags        []string
}

type File struct {
	ID          string
	Name        string
	Content     []byte
	FrontMatter FrontMatter
	Listing     []ListingItem
}

type FrontMatter struct {
	ID          string   `yaml:"id"`
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
			ID:          fm.ID,
			Name:        strings.Replace(dirPath+"/"+strings.Split(file.Name(), ".")[0], "content/", "", 1),
			Content:     []byte(strings.Join(strings.Split(string(buf.Bytes()), "---")[2:], "")),
			FrontMatter: fm,
		}

		if fm.Layout == "listing" {
			listingItem := r.GenerateListing(fileStruct)
			fileStruct.Listing = listingItem
		}

		files = append(files, fileStruct)

		fmt.Println(styles.DebugStyle.Render("Read " + file.Name()))
	}

	return files
}

func (r *Reader) ReadFrontMatter(buf []byte) FrontMatter {
	strBuf := string(buf)
	if !strings.HasPrefix(strBuf, "---") {
		return FrontMatter{}
	}

	fmBuf := []byte(strings.Split(strBuf, "---")[1])
	fm := &FrontMatter{}

	if err := yaml.Unmarshal(fmBuf, fm); err != nil {
		fmt.Println(styles.ErrorStyle.Render("Error reading front matter"))
		os.Exit(1)
	}

	return *fm
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

		buf := bytes.Buffer{}
		if _, err = io.Copy(&buf, f); err != nil {
			fmt.Println(styles.ErrorStyle.Render("Error reading file " + file.Name()))
			os.Exit(1)
		}

		fm := r.ReadFrontMatter(buf.Bytes())
		if fm.Draft {
			continue
		}

		files = append(files, ListingItem{
			ID:          fm.ID,
			Name:        strings.Split(file.Name(), ".")[0],
			URL:         strings.Replace(dirPath+"/"+strings.Split(file.Name(), ".")[0], "content", "", 1),
			Date:        fm.Date,
			Title:       fm.Title,
			Description: fm.Description,
			Tags:        fm.Tags,
		})
	}

	fmt.Println(styles.DebugStyle.Render("Generated listing for " + fileStruct.Name))

	fileStruct.Listing = files
	return files
}
