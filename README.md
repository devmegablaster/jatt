# Jatt

Jatt is a fast and simple static site generator written in Go. It allows you to create and configure static websites with minimal effort. Just set up your site using the `jatt.yaml` configuration file and run the `jatt` command to build it.

## Installation

### Macos

#### Using Homebrew

```bash
brew tap devmegablaster/homebrew-devmegablaster
brew install jatt
```

### Other Platforms

1. Download and install the latest version of Jatt from the [releases](https://github.com/devmegablaster/jatt/releases).
2. Make sure `jatt` is in your system's `PATH`.

## Getting Started

1. Initialize a new Jatt project directory:

   ```bash
   jatt init
   ```

   This will create a new directory with the default `jatt.yaml` configuration file and folder structure.

2. Customize the `jatt.yaml` file as needed:

   ```yaml
   baseURL: <BASE_URL>
   title: <SITE_TITLE>
   author: <AUTHOR_NAME>
   theme: dark # or light
   contentDir: content # directory where your site content is stored
   outputDir: public # directory where the built site will be saved

   nav: #navigation menu config
     title: <TITLE> # navigation menu title
     items:
       - label: Home
         url: /
       - label: Posts
         url: /posts
       - label: About
         url: /about

   home: # home page configuration
     name: <NAME> # name to display on the home page
     description: <DESCRIPTION> # description to display on the home page
     image: <IMAGE_URL> # image to display on the home page
     social:
       - name: github
         url: <GITHUB_URL>
       - name: linkedin
         url: <LINKEDIN_URL>
       - name: <SOCIAL_MEDIA_NAME>
         url: <SOCIAL_MEDIA_URL>
     buttons:
       - name: Posts
         url: /posts
       - name: About
         url: /about
   ```

3. Build your static site:

   ```bash
   jatt
   ```

   The site will be generated in the directory specified by `output_dir`.

## Configuration

The `jatt.yaml` file defines how your site is built. Key fields include:

- `baseUrl`: The base URL for your site.
- `title`: The title of your site.
- `author`: The author of your site.
- `theme`: The theme to use for your site.
- `contentDir`: The directory where your site content is stored.
- `outputDir`: The directory where the built site will be saved.
- `nav`: The navigation menu for your site, with:
  - `title`: The title of the navigation menu.
  - `items`: A list of menu items, each with:
    - `label`: The label for the menu item.
    - `url`: The URL for the menu item.
- `home`: The home page configuration, with:
  - `name`: The name to display on the home page.
  - `description`: The description to display on the home page.
  - `image`: The image to display on the home page.
  - `social`: A list of social media links, each with:
    - `name`: The name of the social media platform.
    - `url`: The URL for the social media platform.
  - `buttons`: A list of buttons to display on the home page, each with:
    - `name`: The name of the button.
    - `url`: The URL for the button.

## Contribution

Contributions are welcome! Whether it’s reporting issues, submitting pull requests, or suggesting features, your help is appreciated.
