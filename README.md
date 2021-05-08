# field-validator
The Field Validator

# Concept:
We define all the checks in config.yaml and then program will check that all the checks are statified in input data ( may be hardcoded , url-response)


If any record (object in js) has any field empty( empty is check. i call it `exist`) (for ex `title`) it will put that record in map and return which check failed on this.

# V1 
Now it only support `exist` check which will make sure that. the field is not empty it has some data. ( field != "")

