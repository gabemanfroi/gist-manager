# Gist Manager

Gist Manager is a command-line tool that simplifies the process of creating Gists on GitHub. With just a single command, you can quickly create Gists using your GitHub API key.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Build from Source](#build-from-source)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Features

- Create Gists on GitHub effortlessly.
- Specify Gist content, description, file extension, and filename.
- Securely authenticate with your GitHub API key.

## Installation

To use Gist Manager, follow these simple installation steps:

1. **Download the Binary**:

   You can download the latest binary for your operating system from the [Releases](https://github.com/gabemanfroi/gist-manager/releases) page.

2. **Install the Binary**:

   Depending on your operating system, you can either place the binary in a directory included in your system's PATH or run it directly from the directory where you saved it.

## Usage

To create a Gist using Gist Manager, open your terminal and run the following command:

```bash
gist-manager create --content "Your Gist content goes here"
```

Replace `"Your Gist content goes here"` with the actual content you want to include in your Gist. You can also specify additional options such as a description, file extension, and filename as needed.

For more detailed usage instructions and examples, refer to the [Command Usage](#command-usage) section in the tool's documentation.

## Configuration

Before using Gist Manager, you need to configure your GitHub API access token by setting it as an environment variable. Follow these steps:

1. **Generate a GitHub Access Token**:

   To create a GitHub access token, follow [GitHub's official guide](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token). Make sure to grant the necessary permissions required for your project.

2. **Set the Environment Variable**:

   Once you have your GitHub access token, you can set it as an environment variable in your terminal. Open a terminal and execute the following command, replacing `YOUR_TOKEN_HERE` with your actual GitHub access token:

   ```bash
   export GITHUB_ACCESS_TOKEN=YOUR_TOKEN_HERE
   ```

   This command sets the `GITHUB_ACCESS_TOKEN` environment variable with your GitHub access token.

3. **Verify the Configuration**:

   To ensure that the environment variable is correctly set, you can run the following command in your terminal:

   ```bash
   echo $GITHUB_ACCESS_TOKEN
   ```

   If the correct token is displayed, your configuration is complete.

## Build from Source

If you prefer to build Gist Manager from source, you can follow these steps:

1. **Clone the Repository**:

   First, clone this GitHub repository to your local machine:

   ```bash
   git clone https://github.com/gabemanfroi/gist-manager.git
   ```

2. **Navigate to the Project Directory**:

   Change your working directory to the project's root folder:

   ```bash
   cd gist-manager
   ```

3. **Build the Binary**:

   Use the `go build` command to build the binary for your platform:

   ```bash
   go build -o gist-manager
   ```

4. **Run the Binary**:

   You can now run the `gist-manager` binary directly from the project folder:

   ```bash
   ./gist-manager create --content "Your Gist content goes here"
   ```

   Replace `"Your Gist content goes here"` with your desired content.

## Dependencies

This project relies on several external packages to function:

- **[Google Go GitHub](https://pkg.go.dev/github.com/google/go-github)**: This project uses the Go client library for the GitHub API provided by Google. You can find more information about this library and its usage in the [official documentation](https://pkg.go.dev/github.com/google/go-github).

- **[Cobra](https://pkg.go.dev/github.com/spf13/cobra)**: Cobra is a popular CLI library for Go, and it is used extensively in this project to define and manage commands. Check out the [Cobra documentation](https://pkg.go.dev/github.com/spf13/cobra) for details on how to work with Cobra commands and flags.

- **[godotenv](https://pkg.go.dev/github.com/joho/godotenv)**: godotenv is used to load environment variables from a `.env` file, which is particularly useful for managing sensitive information like API keys. Learn more about godotenv in the [documentation](https://pkg.go.dev/github.com/joho/godotenv).

- **[survey](https://pkg.go.dev/github.com/AlecAivazis/survey/v2)**: The survey package is used to create interactive command-line prompts and surveys in this tool, enhancing user interaction. Explore the [survey documentation](https://pkg.go.dev/github.com/AlecAivazis/survey/v2) to understand how it can be utilized in your project.

## Contributing

We welcome contributions from the community. If you'd like to contribute to Gist Manager, please follow our [Contribution Guidelines](CONTRIBUTING.md) for details on how to get started.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Thank you for using Gist Manager! If you encounter any issues or have suggestions for improvements, please open an issue on the [GitHub repository](https://github.com/gabemanfroi/gist-manager). We appreciate your feedback!
