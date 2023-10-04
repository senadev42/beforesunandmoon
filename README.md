# BeforeSunAndMoon

BeforeSunAndMoon is a simple CLI tool for terminal stargazing. It allows you to explore random spots in space and provides information about celestial objects using Right Ascension (RA) and Declination (DEC) coordinates.

## Installation

To install BeforeSunAndMoon, you need to have [Go](https://golang.org/doc/install) installed. Once Go is set up, you can install the application using the following command:

```bash
go install github.com/senadev42/beforesunandmoon@latest
```


## Usage

### Explore

Explore a random spot in space and discover what's there:

```bash
beforesunandmoon explore --radius <radius_value>
```

By default, the radius is set to 5000 units. You can specify a custom radius using the `--radius` flag.

### What's Here

Pick a specific spot in space and find out what's there:

```bash
beforesunandmoon whatshere --ra <ra_value> --dec <dec_value> --radius <radius_value>
```

Specify the Right Ascension (RA) and Declination (DEC) values along with an optional radius using the respective flags.

### About

Get information about Right Ascension (RA) and Declination (DEC) coordinates:

```bash
beforesunandmoon about
```

Learn more about how RA and DEC are used as celestial coordinate systems to locate objects in the sky.

## Examples

### Explore a Random Spot

```bash
beforesunandmoon explore
```

### Find Objects at a Specific Location

```bash
beforesunandmoon whatshere --ra 08:23:07.17 --dec -48:29:40.53 --radius 1000
```

## Contributing

If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.