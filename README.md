# AppointyInternTest
This is a backend API for handling to do list.</br>

Functionality of this API-:</br>
</br>

1.Adds a list if the name of the list is given</br>
</br>
2.Deletes a list if the name of the list is given</br>
3.Updates the name of the list corresponding to some list ID.</br>
</br>
4.Adds some item in the particular list</br>
5.Delets the item from the list</br>
6.Updates the completed field of the items</br>
</br>
DATABASE FOR THE API-:</br>
There are two tables </br>
1.list    // for the TODo lists</br>
2.item    // for the items</br>
</br>
list TABLE-:</br>
In the list table there are 2 columns . one for id that is primary key for the list table , another one is name corresponding to the </br>
list name</br>
For every new entry id is different and name of the list is given to the API.</br>
</br>
item TABLE-:</br>
In this table there are 4 columns - id,list_id,value,completed . id is primary key for this table and is corresponding to the item id.</br>
list_id is foreign key from list table to gurranty that there is no item corresponding to unexisting list.</br>
value feild is given to API that is the name of the toDo task or item.</br>
completed is feild to check whether the task has been completed or not.</br>

Checking of the API-:</br>
1.Set the database and create the table as described above.</br>
2.Change in the code section for connection to the database.Change password, host,user as needed.</br>
3.Run the API server.</br>
4.Get a request from POSTMAN.</br>
</br>
you are all set! Bingo!</br>
