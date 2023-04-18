import 'package:flutter/cupertino.dart';
import 'package:app/data/widgets/reminders/get_reminders.dart';
import 'package:app/data/widgets/reminders/create_reminders.dart';


class Reminders extends StatefulWidget {
  const Reminders({Key? key}) : super(key: key);

  @override
  RemindersState createState() => RemindersState();
}

class RemindersState extends State<Reminders> {
  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      child: Padding(
        padding: const EdgeInsets.all(8.0),
        child: Column(
          children: [
            const Text("Reminders"),
            CupertinoButton(
                child: const Text("Show Reminders"),
                onPressed: () => Navigator.push(
                  context,
                  CupertinoPageRoute(
                    builder: (context) => GetReminders(),
                  ),
                ),
            ),
            CupertinoButton(
                child: const Text("Add Reminder"),
                onPressed: () => Navigator.push(
                  context,
                  CupertinoPageRoute(
                    builder: (context) => const CreateReminder(),
                  ),
                ),
            ),
          ],
        ),
      ),
    );
  }
}