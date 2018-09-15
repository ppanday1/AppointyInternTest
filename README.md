# AppointyInternTest
This is a backend API for handling to do list.

Functionality of this API-:


1.Adds a list if the name of the list is given
2.Deletes a list if the name of the list is given
3.Updates the name of the list corresponding to some list ID.

4.Adds some item in the particular list
5.Delets the item from the list
6.Updates the completed field of the items

DATABASE FOR THE API-:
There are two tables 
1.list    // for the TODo lists
2.item    // for the items

list TABLE-:
In the list table there are 2 columns . one for id that is primary key for the list table , another one is name corresponding to the list name
For every new entry id is different and name of the list is given to the API.

item TABLE-:
In this table there are 4 columns - id,list_id,value,completed . id is primary key for this table and is corresponding to the item id.
list_id is foreign key from list table to gurranty that there is no item corresponding to unexisting list.
value feild is given to API that is the name of the toDo task or item.
completed is feild to check whether the task has been completed or not.

Checking of the API-:
1.Set the database and create the table as described above.
2.Change in the code section for connection to the database.Change password, host,user as needed.
3.Run the API server.
4.Get a request from POSTMAN.

you are all set! Bingo!
