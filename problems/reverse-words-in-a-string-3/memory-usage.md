Below is a memory usage analysis based on the code:

a single string with words = X MB

array of bytes created from the string with words ~= X MB

all the words in the array of bytes are reversed with no extra space

a single string from array of bytes, with all words reversed ~= X MB

Total memory used is 3X MB. 3 times the size of the input - single string with words
