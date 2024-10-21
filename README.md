# mbox-splitter

This Go script splits a large mbox file into smaller parts, each containing a maximum number of emails. The split files are named based on the date of the first email in each part. This is useful for managing large email archives and improving performance when processing email data.

## Installation

1. Make sure you have Go installed.
2. Clone this repository:

   ```bash
   git clone https://github.com/VadimOnix/mbox-splitter.git
   cd mbox-splitter
   ```

3. Build the executable:

   ```bash
   go build
   ```

## Usage

```bash
./mbox-splitter <path to mbox file>
```

For example:

```bash
./mbox-splitter /path/to/my/large.mbox
```

This will create multiple `mbox` files in the current directory, named like this:

`mbox_part_<number>_<date>.mbox`

Where:

- `<number>` is the sequential part number (starting from 0).
- `<date>` is the date of the first email in that part, in YYYY_MM_DD format.

## Configuration

- **`maxEmailsPerFile`:** The maximum number of emails per output file. This is controlled by the `maxEmailsPerFile` constant in the `main.go` file and is currently set to 1000. You can modify this value in the source code if needed.
- **`bufferSize`:** The buffer size used for writing to files. This is controlled by the `bufferSize` constant in `main.go` and is set to 4096 bytes. You can adjust this value as needed.

## Example

Let's say you have a large `large.mbox` file. Running:

```bash
./mbox-splitter large.mbox
```

might produce the following output files:

- `mbox_part_0_2024_09_22.mbox`
- `mbox_part_1_2024_09_23.mbox`
- `mbox_part_2_2024_09_24.mbox`
- ... and so on.

## Error Handling

The script includes error handling for file operations and will print error messages to the console if any issues occur.

## Contributing

Contributions are welcome! Please feel free to open issues or submit pull requests.
