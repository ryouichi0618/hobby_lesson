const Example = () => {
  const clickHandler = () => {
    alert("クリックされました")
  }

  const clickHandler2 = () => {
    console.log('ボタンがクリックされました。');
  }

  return (
    <>
      <button onClick={clickHandler}>クリックしてね</button>
      <button onClick={clickHandler2}>クリックしてね</button>
      <button onClick={() => alert("クリックされました") }>クリックしてね</button>
    </>
  );
};

export default Example;
