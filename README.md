# Multi-threaded Program

This program demonstrates the usage of multi-threading in the Go programming language.

## Running the Program

To run the program, follow these steps:

1. Compile the program(from project directory):

   ```sh
   go build .
   ```


2. Run the program with the specified number of threads. Replace `<number of threads>` with your desired thread count:

   ```sh
   ./program <number of threads>
   ```

   For example:

   ```sh
   ./program 4
   ```

The program will demonstrate multi-threading using the specified number of threads.

## Adding Strings to the List

During program execution, you can interactively add strings to the list:

1. Enter a string when prompted and press `Enter` to add it to the list.
2. To finish adding strings and print the list, simply press `Enter` without typing anything.
3. To exit the program, enter `q` and press `Enter`.

## Note


- If you don't provide the number of threads when running, the program will display a usage message.