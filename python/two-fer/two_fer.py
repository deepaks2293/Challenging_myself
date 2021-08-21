#Initalizing with empty variable in function
def two_fer(name=""):
    #checking if name is empty
    if not name:
        name = "you"
    #returing value
    return ("One for {name}, one for me.".format(name=name))

