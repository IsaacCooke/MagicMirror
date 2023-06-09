import { useEffect, useState } from "react";
import { gql, useQuery } from "@apollo/client";
import '../css/Api.scss';
const DisplayPicture = () => {
  const [picture, setPicture] = useState("");
  const [text, setText] = useState("");
  const [photographer, setPhotographer] = useState("");
  const [loading, setLoading] = useState(true);

    const GET_DATA = gql`
        query {
            getNasaApiKey
        }
    `;

  const { loading: queryLoading, error, data } = useQuery(GET_DATA); // Use useQuery directly in the component
  if(error) console.error(error)

  useEffect(() => {
    if (!queryLoading && data) {
      const api_key = data.getNasaApiKey;
      fetch("https://api.nasa.gov/planetary/apod?api_key=" + api_key, {
        headers: {
          Accept: "application/json"
        }
      })
        .then((response) => response.json())
        .then((data) => {
          setPicture(data.url);
          setText(data.explanation);
          setPhotographer(data.title);
          setLoading(false);
        });
    }
  }, [queryLoading, data]);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div className={"api-image-container"}>
      <div className={"api-image-text-background"}>
        <h1 className={"api-image-text"}>{text}</h1>
        <br></br>
        <h2 className={"api-image-text"}><b>{photographer}</b></h2>
      </div>
      <img className={"api-image"} src={picture} alt={"NASA Picture of the Day"} />
    </div>
  );
};

export default DisplayPicture;