document.addEventListener("DOMContentLoaded", () => {
    const taskList = document.getElementById("task-list");
    const taskForm = document.getElementById("task-form");
    const taskInput = document.getElementById("task-input");

    // Sample tasks (in a real app, these would be fetched from the server)
    const tasks = [
        { id: 1, name: "Learn Go" },
        { id: 2, name: "Build a Task Manager" },
    ];

    function renderTasks() {
        taskList.innerHTML = ''; // Clear the list before rendering
        tasks.forEach(task => {
            const li = document.createElement("li");
            li.textContent = task.name;

            const deleteButton = document.createElement("button");
            deleteButton.textContent = "Delete";
            deleteButton.className = "delete-task";
            deleteButton.addEventListener("click", () => {
                deleteTask(task.id);
            });

            li.appendChild(deleteButton);
            taskList.appendChild(li);
        });
    }

    function addTask(taskName) {
        const newTask = {
            id: tasks.length + 1,
            name: taskName
        };
        tasks.push(newTask);
        renderTasks();
    }

    function deleteTask(taskId) {
        const index = tasks.findIndex(task => task.id === taskId);
        if (index !== -1) {
            tasks.splice(index, 1);
            renderTasks();
        }
    }

    taskForm.addEventListener("submit", (e) => {
        e.preventDefault();
        const taskName = taskInput.value.trim();
        if (taskName) {
            addTask(taskName);
            taskInput.value = ''; // Clear the input field
        }
    });

    // Initial rendering of tasks
    renderTasks();
});
