package Algo

//queue with *[]node type
//take graph and key T type to check
//return error if no value T and return nil array of []*node

//push *graph.head
//start a i,v range till length of queue
//if v!=head then pop the from queue
//append v to path array of []*node
//if v.value == key T then return true and v
//else for j,c range v.child
//append queue with c
//
//retuen errorf there was no node with value key , nil
