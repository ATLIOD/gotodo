function TaskList() {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        fetchTasks();
    }, []);

    const fetchTasks = async () => {
        const response = await fetch("http://localhost:8080/tasks");
        const data = await response.json();
        setTasks(data);
    };

    const completeTask = async (id) => {
        try {
            const response = await fetch(
                `http://localhost:8080/tasks/complete?id=${id}`,
                {
                    method: "PATCH",
                },
            );
            if (response.ok) {
                // Refresh task list or update local state
                fetchTasks();
            }
        } catch (error) {
            console.error("Error completing task:", error);
        }
    };

    return (
        <div>
            {tasks.map((task) => (
                <div key={task.id}>
                    <span>{task.name}</span>
                    {!task.complete && (
                        <button onClick={() => completeTask(task.id)}>
                            Complete
                        </button>
                    )}
                </div>
            ))}
        </div>
    );
}
