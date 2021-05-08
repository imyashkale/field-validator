# field-validator
The Field Validator

# Concept:
We define all the checks in config.yaml and then program will check that all the checks are statified in input data ( may be hardcoded , url-response)


If any record (object in js) has any field empty( empty is check. i call it `exist`) (for ex `title`) it will put that record in map and return which check failed on this.

# V1 
Now it only support `exist` check which will make sure that. the field is not empty it has some data. ( field != "")

# v2
You can pass the data to the validator function in main function. And add field in which you want to check for existance ( not empty )
In v1 there is need to define the data stucture in front. now its flexible'

## exist (yes/no)
yes -> field should not be empty
no -> field should be empty

## min -check
len of field value in case of string.
compare with value if its int.

## max -check
same applicable to the max

for example checkout config.yaml

