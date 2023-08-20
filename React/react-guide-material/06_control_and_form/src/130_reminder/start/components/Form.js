import { useState } from "react";

const Form = ({ createTodo }) => {
    const [enteredTodo, setEnteredTodo] = useState("");

    const addTodo = (e) => {
        e.preventDefault();
        const newContent = enteredTodo.trim();
        const inputVal = {
            id: Math.floor(Math.random() * 1e5),
            content: newContent,
        };
        if (newContent === "") {
            alert("値を設定してから追加ボタンを押してください");
            setEnteredTodo("");
            return
        }
        createTodo(inputVal);
        setEnteredTodo("");
    }
    return (
        <>
        <form onSubmit={addTodo}>
            <input type="text" onChange={e => setEnteredTodo(e.target.value)} value={enteredTodo}/>
            <button>追加</button>
        </form>
        </>
    );
}

export default Form;