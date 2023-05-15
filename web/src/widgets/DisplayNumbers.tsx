import {useState, useEffect} from "react";
import '../css/Api.scss';

const DisplayNumbers = () => {
  const [state, refreshState] = useState(false);
  const [numbers, setNumbers] = useState("");

  useEffect(() => {
    setInterval(() => {
      refreshState(!state);
    }, 150000);

    const fetchNumberFact = async () => {
      try {
        const response = await fetch("http://numbersapi.com/random/trivia");
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        const data = await response.text();
        setNumbers(data);
      } catch (error) {
        console.error(error);
      }
    }

    fetchNumberFact();
  }, [state]);

  return (
    <div className={"api-container"}>
      <h1 className={"api-text"}>{numbers}</h1>
    </div>
  );
}

export default DisplayNumbers;