import {gql, useQuery} from "@apollo/client";
import Flashcard, {FlashcardProps} from "../models/Flashcard";
import React, {useEffect, useState} from "react";

const DisplayFlashcards = () => {
    const GET_DATA = gql`
        query {
            getAllFlashcards {
                ID
                Term
                Definition
            }
        }
    `;

  const { loading, error, data } = useQuery(GET_DATA);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error :(</p>;

  const flashcards: FlashcardProps[] = data.getAllFlashcards;
  const shuffledFlashcards = shuffleArray(flashcards);

  return (
    <>
      <FlashcardCycle flashcards={shuffledFlashcards} />
    </>
  );
}

const FlashcardCycle: React.FC<{ flashcards: FlashcardProps[] }> = ({ flashcards }) => {
  const [currentFlashcardIndex, setCurrentFlashcardIndex] = useState(0);

  useEffect(() => {
    const intervalId = setInterval(() => {
      setCurrentFlashcardIndex((currentFlashcardIndex + 1) % flashcards.length);
    }, 30000);

    return () => clearInterval(intervalId);
  }, [flashcards, currentFlashcardIndex]);

  if (flashcards.length === 0) return <></>;
  return (
    <>
      <Flashcard ID={flashcards[currentFlashcardIndex].ID} Term={flashcards[currentFlashcardIndex].Term} Definition={flashcards[currentFlashcardIndex].Definition} />
    </>
  );
};

// A function to take in an array and shuffle it and return a new version of the array
function shuffleArray<T>(array: T[]): T[] {
  const shuffledArray = [...array];
  for (let i = shuffledArray.length - 1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i + 1));
    [shuffledArray[i], shuffledArray[j]] = [shuffledArray[j], shuffledArray[i]];
  }
  return shuffledArray;
}


export default DisplayFlashcards;