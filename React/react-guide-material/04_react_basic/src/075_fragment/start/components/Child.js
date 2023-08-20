import "./Child.css";
import { Fragment } from "react"

const Child = () => {
  return (
    <Fragment>
    <div className="component">
      <h3>Hello Component</h3>
    </div>
    <h1>Fragment</h1>
    <p>
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Obcaecati repellat dolor doloribus iure consequatur soluta? Optio corrupti ratione suscipit recusandae eius perspiciatis illo corporis? Aliquam nam repellendus quos expedita est?
    </p>
    </Fragment>
  );
};

export default Child;
