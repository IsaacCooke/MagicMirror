import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';

class CreateFlashcards extends StatefulWidget {
  @override
  CreateFlashcardsState createState() => CreateFlashcardsState();
}

class CreateFlashcardsState extends State<CreateFlashcards>{
  Constants constants = Constants();

  final _termController = TextEditingController();
  final _definitionController = TextEditingController();

  void _submitForm() async {
    final _client = constants.client;

    final String _term = _termController.text;
    final String _definition = _definitionController.text;

    const mutation = """
    mutation createFlashcard(\$term: String!, \$definition: String!) {
      createFlashcard(term: \$term, definition: \$definition) {
        ID
        Term
        Definition
      }
    }
    """;

    final options = MutationOptions(
      document: gql(mutation),
      variables: {
        "term": _term,
        "definition": _definition,
      },
    );

    final result = await _client.value.mutate(options);

    if(result.hasException) {
      throw Exception(result.exception.toString());
    }
  }

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("Create Flashcards"),
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
                    prefix: const Text("Term"),
                    child: CupertinoTextFormFieldRow(
                      controller: _termController,
                      placeholder: "Term",
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return "Please enter a term";
                        }
                        return null;
                      },
                    ),
                  ),
                  CupertinoFormRow(
                    prefix: const Text("Definition"),
                    child: CupertinoTextFormFieldRow(
                      controller: _definitionController,
                      placeholder: "Definition",
                      validator: (value) {
                        if (value == null || value.isEmpty) {
                          return "Please enter a definition";
                        }
                        return null;
                      },
                    ),
                  ),
                  CupertinoButton(
                    onPressed: () {
                      _submitForm();
                      Navigator.pop(context);
                    },
                    child: const Text("Submit"),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}