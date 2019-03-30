# greedyGame
bid and auction

for adding or deleting any information we need the  user  should be admin .
for that we need to maintain an user table 
where there will be two kind of user 
    simple user and admin 
    that table will contain information related to user

for tracking the adSlot info

we need ad_slot table 
which will hold 
adslot_id, name , current_holder(that will be a registred user),
holding_duration, created_at, updated_at,price ,status('free',"alloted)

bid table which will hold 
adSlotID, objectID, price , any other related information as per requirement.


for running this file 
clone in the src folder of gopath

go to launch folder
 use :- go get
before that you need to create an database and 
for this problem two table ad_slot  and bid
info is send over mail.
 now :- go run main.go
 using postman you can check the exposed api.