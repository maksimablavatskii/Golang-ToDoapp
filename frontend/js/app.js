import { taskListComponent } from "./components/taskList.js";

const formContainer =
    document.querySelector("#task-form-container");

const listContainer =
    document.querySelector("#task-list-container");

formContainer.innerHTML =
    taskFormComponent();

listContainer.innerHTML =
    taskListComponent();