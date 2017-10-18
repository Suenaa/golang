## Selpg
### Usage
　　selpg -s sNumber -e eNumber -l lnumber    

usage of selpg     
`-s` int      
　　the starting index of page, it should be a positive number. The default value is -1.         
`-e` int       
　　the ending index of page, it should be a positive number,too. The default value is -1.          
`-l` int      
　　the number of lines of each page           
`-d` string         
　　printfile name       
`-f` bool           
　　whether seperate pages          

### Examples        
file.txt           
line 1            
line 2          
line 3         
line 4          
line 5            
line 6              
line 7              
line 8              
line 9              
line 10              


`$ selpg -s 2 -e 2 -l 2 file.txt`                

output:              
line 3              
line 4              

`$ selpg -s 2 -e 2 -l 2 -d newFile.txt file.txt`              
output:              
successful write to newFile.txt              

Then `$ selpg -s 1 -e 10 newFile.txt`              
output:              
line 3              
line 4              

`$ selpg -s 2 -e 1 file.txt`              
output:              
start should be less than end              

`$ selpg -s 0 -e 1 file.txt`              
output:              
start should be a positive number              

`$ selpg -s 1 -e 2 -l -3 file.txt`              
output:              
length should be a positive number              

`$ selpg -s 1 -e 2 file`              
output:              
An error occurred on opening the file              

`$ selpg -s 1 -e 2 -d newFile.txt`              
output:              
An error occurred on opening the file              
