class CommitLog:
    def __init__(self, filename):
        self.filename = filename
        open(self.filename, 'a').close()  # ensure file exists

    def append(self, message):
        with open(self.filename, 'a') as f:
            f.write(message + '\n')

    def read_from(self, offset):
        with open(self.filename, 'r') as f:
            lines = f.readlines()
            return lines[offset:], len(lines)


