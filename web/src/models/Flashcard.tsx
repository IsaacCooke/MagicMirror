import '../css/Flashcard.scss';
export interface FlashcardProps {
  ID: string;
  Term: string;
  Definition: string;
}

const Flashcard = (props: FlashcardProps) => {
  return (
    <>
    <div className={"flashcard"}>
      <h1>{props.Term}</h1>
      <p>{props.Definition}</p>
    </div>
    </>
  )
}

export default Flashcard;