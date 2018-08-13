// интерфейс

struct thread { int id; int cpu_usage; struct thread *next; };
struct thread *newThread(int id);

struct queue { int timeslice; struct thread *head, *tail; };
struct queue *newQueue(int timeslice);

void enQueue(struct queue *q, struct thread *t);
struct thread *deQueue(struct queue *q);
struct thread *peek(struct queue *q);
struct thread *find(struct queue *q, int thread_id);

struct queue *process;
struct queue *blocked;

// условие и решение

/**
 * Функция будет вызвана перед каждым тестом, если вы
 * используете глобальные и/или статические переменные
 * не полагайтесь на то, что они заполнены 0 - в них
 * могут храниться значения оставшиеся от предыдущих
 * тестов.
 *
 * timeslice - квант времени, который нужно использовать.
 * Поток смещается с CPU, если пока он занимал CPU функция
 * timer_tick была вызвана timeslice раз.
 **/
void scheduler_setup(int timeslice)
{
    process = newQueue(timeslice);
    blocked = newQueue(0);
}

/**
 * Функция вызывается, когда создается новый поток управления.
 * thread_id - идентификатор этого потока и гарантируется, что
 * никакие два потока не могут иметь одинаковый идентификатор.
 **/
void new_thread(int thread_id)
{
    enQueue(process, newThread(thread_id));
}

/**
 * Функция вызывается, когда поток, исполняющийся на CPU,
 * завершается. Завершится может только поток, который сейчас
 * исполняется, поэтому thread_id не передается. CPU должен
 * быть отдан другому потоку, если есть активный
 * (незаблокированный и незавершившийся) поток.
 **/
void exit_thread()
{
    deQueue(process);
}

/**
 * Функция вызывается, когда поток, исполняющийся на CPU,
 * блокируется. Заблокироваться может только поток, который
 * сейчас исполняется, поэтому thread_id не передается. CPU
 * должен быть отдан другому активному потоку, если таковой
 * имеется.
 **/
void block_thread()
{
    enQueue(blocked, deQueue(process));
}

/**
 * Функция вызывается, когда один из заблокированных потоков
 * разблокируется. Гарантируется, что thread_id - идентификатор
 * ранее заблокированного потока.
 **/
void wake_thread(int thread_id)
{
    enQueue(process, find(blocked, thread_id));
}

/**
 * Ваш таймер. Вызывается каждый раз, когда проходит единица
 * времени.
 **/
void timer_tick()
{
    struct thread *current = peek(process);
    if (current == NULL) {
        return;
    }
    current->cpu_usage += 1;
    if (current->cpu_usage >= process->timeslice) {
        enQueue(process, deQueue(process));
    }
}

/**
 * Функция должна возвращать идентификатор потока, который в
 * данный момент занимает CPU, или -1 если такого потока нет.
 * Единственная ситуация, когда функция может вернуть -1, это
 * когда нет ни одного активного потока (все созданные потоки
 * либо уже завершены, либо заблокированы).
 **/
int current_thread()
{
    struct thread *current = peek(process);
    if (current == NULL) {
        return -1;
    }
    return current->id;
}

// реализация

struct thread *newThread(int id) {
    struct thread *node = (struct thread*)malloc(sizeof(struct thread));
    node->id = id;
    node->next = NULL;
    return node;
}

struct queue *newQueue(int timeslice) {
    struct queue *q = (struct queue*)malloc(sizeof(struct queue));
    q->timeslice = timeslice;
    q->head = q->tail = NULL;
    return q;
}

void enQueue(struct queue *q, struct thread *t) {
    if (t == NULL) {
        return;
    }
    t->cpu_usage = 0;
    t->next = NULL;
    if (q->tail == NULL) {
        q->head = q->tail = t;
        return;
    }
    q->tail->next = t;
    q->tail = t;
}

struct thread *deQueue(struct queue *q) {
    if (q->head == NULL) {
        return NULL;
    }
    struct thread *t = q->head;
    q->head = t->next;
    if (q->head == NULL) {
        q->tail = NULL;
    }
    return t;
}

struct thread *peek(struct queue *q) {
    return q->head;
}

struct thread *find(struct queue *q, int thread_id) {
    struct thread *prev = NULL;
    struct thread *next = peek(q);
    while (next != NULL) {
        if (next->id == thread_id) {
            if (next == q->head) {
                q->head = next->next;
                if (q->head == NULL) {
                    q->tail = NULL;
                }
            } else if (next == q->tail) {
                q->tail = prev;
                prev->next = NULL;
            } else {
                prev->next = next->next;
            }
            return next;
        }
        prev = next;
        next = next->next;
    }
    return NULL;
}
