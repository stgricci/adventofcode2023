import re
if __name__ == "__main__":
    numbers = []
    # read a file line by line
    with open('puzzle_input.txt', 'r') as f:
        for line in f:
            print(line, end='')

            # replace all instances of spelled out numbers with digits
            line = line.replace('zero', 'z0o')
            line = line.replace('one', 'o1e')
            line = line.replace('two', 't2o')
            line = line.replace('three', 't3e')
            line = line.replace('four', 'f4r')
            line = line.replace('five', 'f5e')
            line = line.replace('six', 's6x')
            line = line.replace('seven', 's7n')
            line = line.replace('eight', 'e8t')
            line = line.replace('nine', 'n9e')
            print(line, end='')


            # collect all the numerical digits from the line
            line_numbers = re.findall(r'\d{1}', line)
            print(line_numbers)

            # combined first and last digit
            first_and_last = line_numbers[0] + '' + line_numbers[-1]
            print(first_and_last)
            numbers.append(int(first_and_last))
    
    x = 0
    for n in numbers:
        x += n
    print(x)