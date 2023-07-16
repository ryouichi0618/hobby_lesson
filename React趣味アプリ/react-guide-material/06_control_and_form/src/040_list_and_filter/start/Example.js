import { useState } from "react";

const animals = ["Dog", "Cat", "Rat"];

const Example = () => {
  const [filerVal, setFilterVal] = useState("");
  return (
    <>
      <h3>配列のフィルター</h3>
      <input type="text" value={filerVal} onChange={(e) => {
        setFilterVal(e.target.value);
      }}/>
      <ul>
        {animals
          .map((animal) => (
          <li>{animal}</li>
        ))}
      </ul>
    </>
  );
};

export default Example;
