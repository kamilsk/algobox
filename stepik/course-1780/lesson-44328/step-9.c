void *heap;

struct blob { struct blob *next, *prev; std::size_t size; };

// Эта функция будет вызвана перед тем как вызывать myalloc и myfree
// используйте ее чтобы инициализировать ваш аллокатор перед началом
// работы.
//
// buf - указатель на участок логической памяти, который ваш аллокатор
//       должен распределять, все возвращаемые указатели должны быть
//       либо равны NULL, либо быть из этого участка памяти
// size - размер участка памяти, на который указывает buf
void mysetup(void *buf, std::size_t size)
{
    heap = buf;
    blob *head = (blob*)heap;
    head->next = head->prev = NULL;
    head->size = size;
}

// Функция аллокации
void *myalloc(std::size_t size)
{
    blob *head = (blob*)heap;
    blob **indirect = &head;
    while (*indirect != NULL) {
        if ((*indirect)->size >= size + 2*sizeof(blob)) {
            void *ptr = (void*)(*indirect) + (*indirect)->size - size;
            blob *next = (blob*)((void*)ptr - sizeof(blob));
            next->next = NULL;
            next->prev = *indirect;
            next->size = size + sizeof(blob);
            (*indirect)->next = next;
            (*indirect)->size -= next->size;
            return ptr;
        }
        indirect = &(*indirect)->next;
    }
    return NULL;
}

// Функция освобождения
void myfree(void *ptr)
{
    blob *current = (blob*)((void*)ptr - sizeof(blob));
    current->prev->next = current->next;
    current->prev->size += current->size;
    if (current->next != NULL) {
        current->next->prev = current->prev;
    }
    return;
}
