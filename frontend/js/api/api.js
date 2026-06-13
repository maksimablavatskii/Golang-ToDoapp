const API_URL = "http://localhost:5050/api/v1";

export async function getTasks() {
    const response = await fetch(`${API_URL}/tasks`);

    if (!response.ok) {
        throw new Error("Ошибка загрузки задач");
    }

    return response.json();
}

export async function getStatistics() {
    const response = await fetch(`${API_URL}/statistics`)
    
    if (!response.ok) {
        throw new Error("Ошибка загрузки статистики")
    }

    return response.json()
}

export async function createTask(task) {
    const response = await fetch(`${API_URL}/tasks`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(task)
    });

    if (!response.ok) {
        throw new Error(`HTTP error: ${response.status}`);
    }

    return await response.json();
}