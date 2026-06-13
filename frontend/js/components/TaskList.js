import { getTasks } from "../api/api"

export function taskListComponent() {
    return (`
<div class='tasklist-container'>
    ${console.log(getTasks())}
</div>
`)
}
