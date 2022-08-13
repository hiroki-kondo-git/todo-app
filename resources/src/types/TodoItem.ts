export enum Status { all, todo, doing, done }

export interface TodoItem {
	id: number
	Content: string
	status: Status.todo | Status.doing | Status.done
}