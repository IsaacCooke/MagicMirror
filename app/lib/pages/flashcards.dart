import 'package:flutter/cupertino.dart';

class Flashcards extends StatefulWidget {
  const Flashcards({Key? key}) : super(key: key);

  @override
  FlashcardsState createState() => FlashcardsState();
}

class FlashcardsState extends State<Flashcards> {
  @override
  Widget build(BuildContext context) {
    return const CupertinoPageScaffold(
      child: Center(
        child: Text("Reminders"),
      ),
    );
  }
}