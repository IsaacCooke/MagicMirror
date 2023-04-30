import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/reminder.dart';

class DeleteReminders extends StatefulWidget {
  const DeleteReminders({Key? key}) : super(key: key);

  @override
  DeleteRemindersState createState() => DeleteRemindersState();
}

class DeleteRemindersState extends State<DeleteReminders> {

  static const String query = """
  query {
    getAllReminders {
      ID
      IsDone
      Title
      Repeat
      Description
      DueDate
    }
  }
  """;

  void _deleteReminder(int id) async {
    final client = constants.client;

    const String mutation = """
    mutation deleteReminder(\$id: Int!) {
      deleteReminder(id: \$id) {
        ID
      }
    }
  """;

    final options = MutationOptions(
      document: gql(mutation),
      variables: {
        "id": id,
      },
    );

    final result = await client.value.mutate(options);

    if(result.hasException) {
      throw Exception(result.exception.toString());
    }
  }

  final Constants constants = Constants();

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("reminders"),
      ),
      child: GraphQLProvider(
        client: constants.client,
        child: Query(
          options: QueryOptions(
            document: gql(query),
          ),
          builder: (result, {fetchMore, refetch}){
            if (result.isLoading){
              return const Center(
                child: CupertinoActivityIndicator(),
              );
            }
            if (result.data == null){
              return const Center(
                child: Text("No Data"),
              );
            }
            final reminders = result.data!["getAllReminders"] as List<dynamic>;
            return ListView.builder(
              itemCount: reminders.length,
              itemBuilder: (context, index){
                final data = reminders[index];
                final int id = data['ID'];
                final String title = data['Title'];
                final String description = data['Description'];
                final String dueDate = data['DueDate'];
                final bool isDone = data['IsDone'];
                final bool repeat = data['Repeat'];
                return Row(
                  mainAxisSize: MainAxisSize.min,
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Flexible(
                      child: ReminderWidget(
                        id: id,
                        title: title,
                        description: description,
                        dueDate: dueDate,
                        isDone: isDone,
                        repeat: repeat,
                      ),
                    ),
                    CupertinoButton(
                      child: const Icon(CupertinoIcons.delete),
                      onPressed: () {
                        _deleteReminder(id);
                        refetch!();
                      },
                    ),
                  ],
                );
              },
            );
          },
        ),
      ),
    );
  }
}