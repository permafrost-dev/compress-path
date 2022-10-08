<p align="center">
<img style="height: 75px; width: auto" src="https://user-images.githubusercontent.com/5508707/194621868-c41519e6-4499-435d-a559-6f5719f7ed98.png" alt="logo" />
</p>

# compress-path

[![Release](https://img.shields.io/github/release/permafrost-dev/compress-path.svg)](https://github.com/permafrost-dev/compress-path/releases/latest) [![Go Report Card](https://goreportcard.com/badge/github.com/permafrost-dev/compress-path)](https://goreportcard.com/report/github.com/permafrost-dev/compress-path)

---

Compress and abbreviate the parts of a pathname into a shorter overall string; for use within a Zsh prompt. Adds some minor styling, such as bold text for the last path section (the current directory).

![image](https://user-images.githubusercontent.com/5508707/194622385-96e6e616-5cfa-4305-9e11-311f28f64432.png)

## Setup

Download and extract `compress-path` to your `$PATH`. Then, add the following to your `.zshrc`:

```bash
export PS1='$(compress-path) >'
```

If you're using `oh-my-zsh`, edit your `~/.oh-my-zsh/your-theme.zsh-theme` and modify the `prompt_dir` (or similar) function to look something like this:

```bash
prompt_dir() {
  prompt_segment blue $CURRENT_FG $(compress-path)
}
```

## Building from source

`compress-path` prefers `task` for running tasks - see [taskfile.dev](https://taskfile.dev) for more information. If you don't have `task` installed, you can use `make` instead.

```bash
task build
# or
make build
```

---

## Changelog

Please see [CHANGELOG](CHANGELOG.md) for more information on what has changed recently.

## Contributing

Please see [CONTRIBUTING](.github/CONTRIBUTING.md) for details.

## Security Vulnerabilities

Please review [our security policy](../../security/policy) on how to report security vulnerabilities.

## Credits

- [Patrick Organ](https://github.com/patinthehat)
- [All Contributors](../../contributors)

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.
