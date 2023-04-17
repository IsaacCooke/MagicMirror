import 'package:flutter/cupertino.dart';

@immutable
class ReminderWidget extends StatelessWidget {
  final int id;
  final String title;
  final String description;
  final String dueDate;
  final bool isDone;
  final bool repeat;

  const ReminderWidget({
    Key? key,
    required this.id,
    required this.title,
    required this.description,
    required this.dueDate,
    required this.isDone,
    required this.repeat,
  }) :super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
        padding: const EdgeInsets.all(8.0),
        child: Row (
            children: [
              Expanded (
                flex: 1,
                child: title != null
                    ? Text(title)
                    : const Text('No data'),
              ),
              Expanded (
                  flex: 2,
                  child: Column (
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text(id.toString()),
                      Text(title),
                      Text(description),
                      Text(dueDate),
                      Text(isDone.toString()),
                      Text(repeat.toString()),
                    ],
                  )
              )
            ]
        )
    );
  }
}