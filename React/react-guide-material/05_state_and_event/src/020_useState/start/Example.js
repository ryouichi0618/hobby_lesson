import { useState } from "react";

const Example = () => {
  let [val, setFn] = useState()

  return (
    <>
      <input
        type="text"
        onChange={(e) => {
          setFn(e.target.value)
        }}
      /> = {val}
    </>
  );
};

export default Example;
