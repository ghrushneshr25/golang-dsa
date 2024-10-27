/*
Problem-45 	
Given a linked list consists of data, a next pointer and also a random pointer which points to a random node of the list. 
Give an algorithm for cloning the list. 

	
Solution: 	To clone a linked list with random pointers, the idea is to maintain a hash table for storing the mappings from a original linked list 
			node to its clone. For each node in the original linked list, we create a new node with same data and set its next pointers. While doing so, we 
			also create a mapping from the original node to the duplicate node in the hash table. Finally, we traverse the original linked list again and 
			update random pointers of the duplicate nodes using the hash table. 


Algorithm: 
	- Scan the original list and for each node X, create a new node Y with data of X, then store the pair (X,Y) in hash table using X as a 
	key. Note that during this scan set Y → NEXT  with X → NEXT and Y → RANDOM to NULL and we will fix it in the next scan. Now 
	for each node X in the original list we have a copy Y stored in our hash table. 
	- To update the random pointers, read random pointer of node X from original linked list and get the corresponding random node 
	in cloned list from hash table created in previous step. Assign random pointer of node X from cloned list to corresponding node 
	we got. 
*/

package linkedlist