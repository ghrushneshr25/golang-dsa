package linkedlist

type RemoveDuplicatesFromUnorderedList struct {
    head   *SingleLinkedList
    Result []int
}

func NewRemoveDuplicatesFromUnorderedList(val []int) *RemoveDuplicatesFromUnorderedList {
    return &RemoveDuplicatesFromUnorderedList{
        head: ConvertArrayToSingleLinkedList(val),
    }
}

func (obj *RemoveDuplicatesFromUnorderedList) RemoveDuplicates() {
    presentMap := make(map[int]bool)

    if obj.head == nil {
        return
    }

    current := obj.head
    var prev *SingleLinkedList

    for current != nil {
        if _, ok := presentMap[current.GetData()]; ok {
            prev.SetNext(current.GetNext())
        } else {
            presentMap[current.GetData()] = true
            prev = current
        }
        current = current.GetNext()
    }
}

func (obj *RemoveDuplicatesFromUnorderedList) Execute() {
    obj.RemoveDuplicates()
    obj.Result = ConvertSingleLinkedListToArray(obj.head)
}