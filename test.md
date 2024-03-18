//Current solution : 

map[A][B] = 100 
nap[B][A] = 50 
map[A][C] = 150


//Alternate : 
[]string: ["A:B:100", "B:A:50", "A:C:150"] // This will be stored in db

[OWES_TABLE] :
|owed_id->user_id(sec_key)|ower_id->user_id(sec key)|amount|group_id->group_id(sec_key)
[GROUP_TABLE]: 
|group_id|user_id->user_id(sec_id)|
[USER_TABLE]: 
|user_id(p_key)|upi|email(p_key)|


for i in array : 
    // Considering delimited values : 
    if map[delimit2+delimit1] == 0:
        map[delimit1+delimit2] += delimit3
    if delimit3 < map[delimit2+delimit1] : 
        map[delimit2+delimit1] == delimit3 
    else : 

for i in maps.key : 
    if i[0] == "A"
        print(maps[i])
        


