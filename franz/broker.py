from storage import CommitLog

class Broker:
    def __init__(self):
        self.topics = {}

    def create_topic(self, name, partitions=1):
        if name not in self.topics:
            self.topics[name] = [CommitLog(f"{name}_part{i}.log") for i in range(partitions)]

    def produce(self, topic, message, partition=0):
        self.topics[topic][partition].append(message)

    def consume(self, topic, partition=0, offset=0):
        return self.topics[topic][partition].read_from(offset)


