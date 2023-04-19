import 'package:flutter/cupertino.dart';

import 'package:app/data/widgets/flashcards/create_flashcards.dart';
import 'package:app/data/widgets/flashcards/get_flashcards.dart';

class Flashcards extends StatefulWidget {
  const Flashcards({Key? key}) : super(key: key);

  @override
  FlashcardsState createState() => FlashcardsState();
}

class FlashcardsState extends State<Flashcards> {
  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      child: Padding(
        padding: EdgeInsets.all(8.0),
        child: Column(
          children: [
            CupertinoButton(
              child: Text("Show Flashcards"),
              onPressed: () => Navigator.push(
                context,
                CupertinoPageRoute(
                  builder: (context) => GetFlashcards(),
                ),
              ),
            ),
            CupertinoButton(
              child: Text("Add Flashcard"),
              onPressed: () => Navigator.push(
                context,
                CupertinoPageRoute(
                  builder: (context) => CreateFlashcards(),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}