import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';

class CreateNotes extends StatefulWidget {
  @override
  CreateNotesState createState() => CreateNotesState();
}

class CreateNotesState extends State<CreateNotes> {
  Constants constants = Constants();

  final _contentController = TextEditingController();

  void _submitForm() async {
    final _client = constants.client;

    final String _content = _contentController.text;

    const mutation = """
    mutation createNote(\$content: String!) {
      createNote(content: \$content) {
        ID
        Content
      }
    }
    """;

    final options = MutationOptions(
      document: gql(mutation),
      variables: {
        "content": _content,
      },
    );

    final result = await _client.value.mutate(options);

    if(result.hasException) {
      throw Exception(result.exception.toString());
    }
  }

  @override
  Widget build(BuildContext context){
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("Create Notes"),
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
                    prefix: const Text("Content"),
                    child: CupertinoTextFormFieldRow(
                      controller: _contentController,
                      placeholder: "Content",
                      onSaved: (value) {
                        _contentController.text = value!;
                      },
                    ),
                  ),
                  CupertinoButton(
                    onPressed: () {
                      _submitForm();
                    },
                    child: const Text("Submit"),
                  ),
                ],
              )
            ),
          ],
        ),
      ),
    );
  }
}