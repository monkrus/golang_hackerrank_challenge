/* 13. 
Complete the Difference class by writing the following:

A class constructor that takes an array of integers as a parameter and saves it to the elements instance variable.
A computeDifference method that finds the maximum absolute difference between any 2 numbers in N and stores it in the maximumDifference instance variable.*/

class Difference:
    def __init__(self, a):
        self.__elements = a

	# Add your code here

        self.maximumDifference = 0

    def computeDifference(self):
        
        for i in self.__elements:
            for j in self.__elements:

                diff = abs(i - j)
                
                if diff > self.maximumDifference:
                    self.maximumDifference = diff
# End of Difference class

_ = input()
a = [int(e) for e in input().split(' ')]

d = Difference(a)
d.computeDifference()

print(d.maximumDifference)
