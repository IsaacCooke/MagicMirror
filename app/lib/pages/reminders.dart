import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Reminders extends StatefulWidget {
  const Reminders({Key? key}) : super(key: key);

  @override
  RemindersState createState() => RemindersState();
}

class RemindersState extends State<Reminders> {
  @override
  Widget build(BuildContext context) {
    return const CupertinoPageScaffold(
      child: Center(
        child: Text("Reminders"),
      ),
    );
  }
}