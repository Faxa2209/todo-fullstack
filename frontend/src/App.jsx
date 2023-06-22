import React, { useState, useEffect } from "react";
import axios from "axios";

const TodoApp = () => {
  const [todos, setTodos] = useState([]);
  const [newTodo, setNewTodo] = useState("");
  const [newTodoBody, setNewTodoBody] = useState("");

  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    try {
      const response = await axios.get("http://localhost:8080/api/todos");
      setTodos(response.data);
    } catch (error) {
      console.error("Failed to fetch todos:", error);
    }
  };

  const addTodo = async () => {
    try {
      const todo = {
        title: newTodo,
        done: false,
        body: newTodoBody,
      };

      const response = await axios.post("http://localhost:8080/api/todos", todo);
      setTodos(response.data);
      setNewTodo("");
      setNewTodoBody("");
    } catch (error) {
      console.error("Failed to add todo:", error);
    }
  };

  const toggleTodo = async (id) => {
    try {
      const response = await axios.patch(`http://localhost:8080/api/todos/${id}`);
      setTodos(response.data);
    } catch (error) {
      console.error("Failed to toggle todo:", error);
    }
  };

  const deleteTodo = async (id) => {
    try {
      const response = await axios.delete(`http://localhost:8080/api/todos/${id}`);
      setTodos(response.data);
    } catch (error) {
      console.error("Failed to delete todo:", error);
    }
  };

  return (
    <div>
      <h1>Todo App</h1>

      <div>
        <input type="text" value={newTodo} onChange={(e) => setNewTodo(e.target.value)} placeholder="Enter a new todo" />
        <input
          type="text"
          value={newTodoBody}
          onChange={(e) => setNewTodoBody(e.target.value)}
          placeholder="Enter the body text"
        />
        <button onClick={addTodo}>Add Todo</button>
      </div>

      <ul>
        {todos.map((todo) => (
          <li key={todo.id}>
            <strong className={`${todo.done ? "done" : ""}`}>{todo.title}</strong>
            <br />
            {todo.body}
            <br />
            <button onClick={() => toggleTodo(todo.id)}>Mark as Done</button>
            <button onClick={() => deleteTodo(todo.id)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default TodoApp;
