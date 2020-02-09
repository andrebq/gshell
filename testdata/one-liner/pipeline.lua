program(
    pipeline([
        binary_program("echo", argument("hello")),
        binary_program("cat", argument("output.txt"))
    ])
)
