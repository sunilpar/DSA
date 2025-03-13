#include <iostream>
using namespace std;

struct Node
{
    int value;
    Node *next;
};

class Queue
{
    Node *head;
    Node *tail;

public:
    Queue() : head(nullptr), tail(nullptr) {}

    ~Queue()
    {
        while (!isEmpty())
        {
            dequeue();
        }
    }

    void enqueue(int v)
    {
        Node *newNode = new Node{v, nullptr};
        if (tail == nullptr)
        {
            head = tail = newNode;
        }
        else
        {
            tail->next = newNode;
            tail = newNode;
        }
    }

    void dequeue()
    {
        if (head == nullptr)
            return;
        Node *temp = head;
        head = head->next;
        if (head == nullptr)
            tail = nullptr;
        delete temp;
    }

    int peek()
    {
        if (!isEmpty())
            return head->value;
        else
            throw runtime_error("Queue is empty");
    }

    void display()
    {
        if (isEmpty())
        {
            printf("The queue is empty!\n");
            return;
        }
        Node *newNode = head;
        while (newNode != nullptr)
        {
            printf("%d, ", newNode->value);
            newNode = newNode->next;
        }
        printf("\n");
    }

    bool isEmpty()
    {
        return head == nullptr;
    }
};

int main()
{
    Queue q;
    q.enqueue(10);
    q.enqueue(20);
    q.enqueue(30);
    q.display();

    q.dequeue();
    q.dequeue();

    q.display();

    return 0;
}

// **Command to Compile & Run:**
// `gcc -o FifoQueue.exe FifoQueue.cpp -lstdc++ && ./FifoQueue.exe`
