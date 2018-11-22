class Person:
    def __init__(self, first_name: str, last_name: str, id_number: int):
        self.firstName = first_name
        self.lastName = last_name
        self.idNumber = id_number

    def print_person(self):
        print("Name:", self.lastName + ",", self.firstName)
        print("ID:", self.idNumber)


class Student(Person):
    def __init__(self, first_name: str, last_name: str, person_id: int, score_value: list):
        super(Student, self).__init__(first_name, last_name, person_id)
        self.scores = score_value

    def calculate(self):
        average = sum(self.scores) / len(self.scores)
        if 90 <= average <= 100:
            return 'O'
        if 80 <= average:
            return 'E'
        if 70 <= average:
            return 'A'
        if 55 <= average:
            return 'P'
        if 40 <= average:
            return 'D'
        return 'T'
