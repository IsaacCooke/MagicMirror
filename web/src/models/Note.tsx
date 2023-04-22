export interface NoteProps {
  ID: string;
  Content: string;
}

const Note = (props: NoteProps) => {
  return (
    <>
    <div>
      <p>{props.Content}</p>
    </div>
    </>
  )
}

export default Note;