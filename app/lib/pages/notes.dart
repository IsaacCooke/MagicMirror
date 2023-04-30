import 'package:app/data/widgets/notes/create_notes.dart';
import 'package:app/data/widgets/notes/get_notes.dart';
import 'package:flutter/cupertino.dart';

class Notes extends StatefulWidget {
  const Notes({Key? key}) : super(key: key);

  @override
  NotesState createState() => NotesState();
}

class NotesState extends State<Notes>{

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      child: Padding(
        padding: EdgeInsets.all(8.0),
        child: Column(
          children: [
            CupertinoButton(
                child: const Text("Show Notes"),
                onPressed: () => Navigator.push(
                  context,
                  CupertinoPageRoute(
                    builder: (context) => GetNotes(),
                  ),
                ),
            ),
            CupertinoButton(
                child: const Text("Add Note"),
                onPressed: () => Navigator.push(
                  context,
                  CupertinoPageRoute(
                    builder: (context) => CreateNotes(),
                  ),
                ),
            ),
          ],
        ),
      ),
    );
  }
}