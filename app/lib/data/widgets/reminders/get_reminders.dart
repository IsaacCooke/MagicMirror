import 'package:flutter/cupertino.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/reminder.dart';

class GetReminders extends StatelessWidget {
  GetReminders({Key? key}) : super(key: key);

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

  final Constants constants = Constants();

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text('Reminders'),
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
                child: Text('No data'),
              );
            }
            final reminders = result.data!['getAllReminders'] as List<dynamic>;
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
                return ReminderWidget(
                  id: id,
                  title: title,
                  description: description,
                  dueDate: dueDate,
                  isDone: isDone,
                  repeat: repeat,
                );
              },
            );
          },
        ),
      )
    );
  }
}