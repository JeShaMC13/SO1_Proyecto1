#include <linux/fs.h>
#include <linux/init.h>
#include <linux/kernel.h>
#include <linux/list.h>
#include <linux/module.h>
#include <linux/proc_fs.h>
#include <linux/sched.h>
#include <linux/seq_file.h>
#include <linux/slab.h>
#include <linux/string.h>
#include <linux/types.h>


void readProcess(struct seq_file *m, struct task_struct *s)
{
    struct list_head *list;
    struct task_struct *task;
    
    char estado[50];

    switch(s->state){
        case TASK_RUNNING:
	    strcpy(estado,"Ejecucion");
	    break;

	case TASK_STOPPED:
 	    strcpy(estado,"Detenido");
	    break;

	case TASK_INTERRUPTIBLE:
 	    strcpy(estado,"Interrumpible");
	    break;

	case TASK_UNINTERRUPTIBLE:
 	    strcpy(estado,"Ininterrumpible");
	    break;

	case EXIT_ZOMBIE:
 	    strcpy(estado,"Zombi");
	    break;

	default:
	    strcpy(estado, "Desconocido");
    }

    seq_printf(m,"PID: %d\t\tNombre: %s\t\tEstado:%s\n",s->pid, s->comm, estado);
    
    list_for_each(list, &s->children) {
        task = list_entry(list, struct task_struct, sibling);
        readProcess(m, task);
    }
}

static int pstree(struct seq_file *m, void *v)
{
    struct task_struct *parent = current;
    while (parent->pid != 1)
        parent = parent->parent; 
    readProcess(m, parent);
    return 0;
}

static int meminfo_proc_open(struct inode *inode, struct file *file)
{
    return single_open(file, pstree, NULL);
}

static const struct file_operations meminfo_proc_fops = {
    .open       = meminfo_proc_open,
    .read       = seq_read,
    .llseek     = seq_lseek,
    .release    = single_release,
};

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Pedro Garcia/Jenny Montenegro");
MODULE_DESCRIPTION("Modulo de CPU - Sistemas Operativos 1 Junio 2020");

static int __init cpu_201318664_200721968_init(void)
{
	printk(KERN_INFO "Pedro Garcia/Jenny Montenegro \n");
	proc_create("cpu_201318664_200721968", 0, NULL, &meminfo_proc_fops);
	return 0;
}

static void __exit cpu_201318664_200721968_cleanup(void)
{
	remove_proc_entry("cpu_201318664_200721968", NULL);
	printk(KERN_INFO "Sistemas Operativos 1\n");
}

module_init(cpu_201318664_200721968_init);
module_exit(cpu_201318664_200721968_cleanup);
