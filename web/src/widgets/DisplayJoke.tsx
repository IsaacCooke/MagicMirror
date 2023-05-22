import { useState, useEffect } from "react";
import '../css/Api.scss';

const DisplayJoke = () => {
  const [state, refreshState] = useState(false);
  const [joke, setJoke] = useState("");

  useEffect(() => {
    setInterval(() => {
      refreshState(!state);
    }, 30000);

    fetch("https://v2.jokeapi.dev/", {
      headers: {
        Accept: "application/json"
      }
    })
      .then(response => response.json())
      .then(data => setJoke(data.joke));
  }, [state]);


  return (
    <div className={"api-container"}>
      <h1 className={"api-text"}>{joke}</h1>
    </div>
  );
}

export default DisplayJoke;