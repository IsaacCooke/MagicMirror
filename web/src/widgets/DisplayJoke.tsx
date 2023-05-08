import { useState, useEffect } from "react";
import '../css/Joke.scss';

const DisplayJoke = () => {
  const [state, refreshState] = useState(false);
  const [joke, setJoke] = useState("");

  useEffect(() => {
    setInterval(() => {
      refreshState(!state);
    }, 30000);

    fetch("https://icanhazdadjoke.com/", {
      headers: {
        Accept: "application/json"
      }
    })
      .then(response => response.json())
      .then(data => setJoke(data.joke));
  }, [state]);


  return (
    <div className={"joke-container"}>
      <h1 className={"joke-text"}>{joke}</h1>
    </div>
  );
}

export default DisplayJoke;