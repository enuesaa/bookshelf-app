import { DeleteTeapod } from './DeleteTeapod'
import { TeapodDescription } from './TeapodDescription'
import { TeapodSchemas } from './TeapodSchemas'
import { useListTeapods } from '@/lib/api/teapods'
import { Heading, Text, Table } from '@radix-ui/themes'

export const ListTeapods = () => {
  const teapods = useListTeapods()

  return (
    <>
      <Heading as='h3'>Teapods</Heading>
      <Text as='p' size='4' mt='2' mb='6' color='gray'></Text>

      <Table.Root variant='surface'>
        <Table.Header>
          <Table.Row>
            <Table.ColumnHeaderCell width='10%'>Name</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='20%'>Description</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell>Teaboxes</Table.ColumnHeaderCell>
            <Table.ColumnHeaderCell width='10%'></Table.ColumnHeaderCell>
          </Table.Row>
        </Table.Header>
        <Table.Body>
          {teapods !== undefined &&
            teapods.data?.items.map((p, i) => (
              <Table.Row key={i}>
                <Table.RowHeaderCell>{p.name}</Table.RowHeaderCell>
                <Table.Cell>
                  <TeapodDescription name={p.name} />
                </Table.Cell>
                <Table.Cell>
                  <TeapodSchemas name={p.name} />
                </Table.Cell>
                <Table.Cell>
                  <DeleteTeapod name={p.name} />
                </Table.Cell>
              </Table.Row>
            ))}
        </Table.Body>
      </Table.Root>
    </>
  )
}
