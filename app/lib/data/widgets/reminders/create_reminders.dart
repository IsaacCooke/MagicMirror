import 'package:flutter/cupertino.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';

class CreateReminder extends StatefulWidget {
  const CreateReminder({super.key});

  @override
  State<CreateReminder> createState() => CreateReminderState();
}

class CreateReminderState extends State<CreateReminder> {
  Constants constants = Constants();

  final _titleController = TextEditingController();
  final _descriptionController = TextEditingController();
  DateTime _dueDate = DateTime.now();
  bool repeat = false;

  void _showDialog(Widget child) {
    showCupertinoModalPopup<void>(
        context: context,
        builder: (BuildContext context) => Container(
          height: 216,
          padding: const EdgeInsets.only(top: 6.0),
          margin: EdgeInsets.only(
            bottom: MediaQuery.of(context).viewInsets.bottom,
          ),
          color: CupertinoColors.systemBackground.resolveFrom(context),
          child: SafeArea(
            top: false,
            child: child,
          ),
        ),
    );
  }

  void _submitForm() async {
    final _client = GraphQLClient(
      link: HttpLink('https://27d8-146-70-95-126.ngrok-free.app/graphql'),
      cache: GraphQLCache(),
    );

    final title = _titleController.text;
    final description = _descriptionController.text;

    const mutation = '''
    mutation createReminder(\$title: String!, \$description: String!, \$dueDate: String!, \$repeat: Boolean!) {
      createReminder(title: \$title, description: \$description, dueDate: \$dueDate, repeat: \$repeat) {
        ID
      }
    }
  ''';

    final options = MutationOptions(
      document: gql(mutation),
      variables: <String, dynamic>{
        'title': title,
        'description': description,
        'dueDate': _dueDate.toIso8601String(),
        'repeat': repeat,
      },
    );

    final result = await _client.mutate(options);

    if (result.hasException) {
      throw Exception(result.exception);
    }
  }

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("Create Reminder"),
      ),
      child: SafeArea(
        child: Column(
          children: <Widget>[
            Form(
              autovalidateMode: AutovalidateMode.always,
              onChanged: () {
                Form.of(primaryFocus!.context!)!.save();
              },
              child: CupertinoFormSection.insetGrouped(
                children: [
                  CupertinoFormRow(
                    prefix: const Text("Title"),
                    child: CupertinoTextFormFieldRow(
                      controller: _titleController,
                      placeholder: "Title",
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return "Please enter a title";
                        }
                        return null;
                      },
                    ),
                  ),
                  CupertinoFormRow(
                    prefix: const Text("Description"),
                    child: CupertinoTextFormFieldRow(
                      controller: _descriptionController,
                      placeholder: "Description",
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return "Please enter a description";
                        }
                        return null;
                      },
                    ),
                  ),
                  CupertinoButton(
                    onPressed: () => _showDialog(
                      CupertinoDatePicker(
                        mode: CupertinoDatePickerMode.dateAndTime,
                        initialDateTime: _dueDate,
                        onDateTimeChanged: (dateTime) {
                          setState(() {
                            _dueDate = dateTime;
                          });
                        },
                      ),
                    ),
                    child: Text(
                      '${_dueDate.day}-${_dueDate.month}-${_dueDate.year} ${_dueDate.hour}:${_dueDate.minute}'
                    ),
                  ),
                  CupertinoFormRow(
                    prefix: const Text("Repeat"),
                    child: CupertinoSwitch(
                      value: repeat,
                      onChanged: (value) {
                        setState(() {
                          repeat = value;
                        });
                      },
                    ),
                  ),
                  CupertinoButton(
                    onPressed: () {
                      _submitForm();
                      Navigator.pop(context);
                    },
                    child: const Text("Create Reminder"),
                  ),
                ],
              ),
            ),
          ]
        )
      ),
    );
  }
}